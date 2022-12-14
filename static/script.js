// COLORS
// ================================================================================
const TURQUOISE = "rgb(111, 218, 243)"
const BROWN = "#555759"
const LIGHT_GREY = "#f6f6f6"
const PINK = "#c85e7a"

// COLUMN FUNCTIONS
// ================================================================================
const getMarked = (row) => {
    for (const column of row.children) {
        if (column.style.background == TURQUOISE) {
            return column
        }
    }
}

const markColumn = (column) => {
    column.style.background = TURQUOISE
    column.style.color = BROWN
}

const unmarkColumn = (column) => {
    column.style.background = PINK
    column.style.color = LIGHT_GREY
}

// SORT & LOAD WORKOUTS
// ================================================================================
const sortAndLoad = async (row, table, column, workouts) => {
    workouts.reverse()
    const marked = getMarked(row)
    if (column.id != marked.id) {
        unmarkColumn(marked)
        markColumn(column)
    }
    clearTable(table)
    loadWorkouts(table, workouts)
    column.onclick = () => { sortAndLoad(row, table, column, workouts) }
}
const clearTable = (table) => {
    table.innerHTML = ""
} 

// LOAD SORTERS
// ================================================================================
const loadSorters = (row, table, workouts) => {
    // date (comes already sorted)
    const workoutsByDate = workouts.slice()
    const date = document.getElementById("date")
    date.onclick = () => { sortAndLoad(row, table, date, workoutsByDate) }

    // distance
    const distance = document.getElementById("distance")
    const workoutsByDist = workouts.slice()
    workoutsByDist.sort((a, b) => { return a.distance - b.distance })
    distance.onclick = () => { sortAndLoad(row, table, distance, workoutsByDist) }

    // duration
    const duration = document.getElementById("duration")
    const workoutsByDur = workouts.slice()
    workoutsByDur.sort((a, b) => { return a.duration.front - b.duration.front })
    duration.onclick = () => { sortAndLoad(row, table, duration, workoutsByDur) }

    // elevation
    const elevation = document.getElementById("elevation")
    const workoutsByElev = workouts.slice()
    workoutsByElev.sort((a, b) => { return a.elevation - b.elevation })
    elevation.onclick = () => { sortAndLoad(row, table, elevation, workoutsByElev) }

    // pace
    const pace = document.getElementById("pace")
    const workoutsByPace = workouts.slice()
    workoutsByPace.sort((a, b) => { return a.avgPace - b.avgPace })
    pace.onclick = () => { sortAndLoad(row, table, pace, workoutsByPace) }

    // hr
    const hr = document.getElementById("hr")
    const workoutsByHR = workouts.slice()
    workoutsByHR.sort((a, b) => { return a.avgHR - b.avgHR })
    hr.onclick = () => { sortAndLoad(row, table, hr, workoutsByHR) }
}

// LOAD WORKOUTS
// ================================================================================
const loadWorkouts = (table, workouts) => {
    for (const workout of workouts) {
        const tr = document.createElement("tr")
        addRowData(tr, workout.date)
        addRowData(tr, `${workout.distance}km`)
        addRowData(tr, workout.duration.back)
        addRowData(tr, `${workout.elevation}m`)
        addRowData(tr, workout.avgPace)
        addRowData(tr, workout.avgHR)
        table.append(tr)
    }
}
const addRowData = (tr, data) => {
    td = document.createElement("td")
    td.innerHTML = data
    tr.append(td)
}

// ON PAGE LOAD
// ================================================================================
const main = async () => {
    
    // init constants
    const row = document.getElementById("first_row")
    const table = document.getElementById("workouts")
    const date = document.getElementById("date")

    // request workouts
    const response = await fetch("/workouts")
    if (response.ok) {
        const workouts = await response.json()

        // if no workouts
        if (workouts.length == 0) {
            return
        }

        // load workouts in DOM
        loadWorkouts(table, workouts)
        markColumn(date)

        // load sorters
        loadSorters(row, table, workouts)
    }

    else {
        console.log(response.status, response.statusText)
    }
}

main()