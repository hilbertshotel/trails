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
    const date = document.getElementById("date")
    date.onclick = () => { sortAndLoad(row, table, date, workouts) }

    // distance
    const distance = document.getElementById("distance")
    workouts.sort((a, b) => { return a.Distance - b.Distance })
    distance.onclick = () => { sortAndLoad(row, table, distance, workouts) }

    // duration
    const duration = document.getElementById("duration")
    duration.onclick = () => { sortAndLoad(row, table, duration, sortDuration(workouts)) }

    // elevation
    const elevation = document.getElementById("elevation")
    workouts.sort((a, b) => { return a.Elevation - b.Elevation })
    elevation.onclick = () => { sortAndLoad(row, table, elevation, workouts) }

    // pace
    const pace = document.getElementById("pace")
    workouts.sort((a, b) => { return a.AvgPace - b.AvgPace })
    pace.onclick = () => { sortAndLoad(row, table, pace, workouts) }

    // hr
    const hr = document.getElementById("hr")
    workouts.sort((a, b) => { return a.AvgHR - b.AvgHR })
    hr.onclick = () => { sortAndLoad(row, table, hr, workouts) }
}

const sortDuration = (workouts) => {
    workouts.sort((a, b) => {
        

        return a.Elevation - b.Elevation
    })
    return workouts
}

// LOAD WORKOUTS
// ================================================================================
const loadWorkouts = (table, workouts) => {
    for (const workout of workouts) {
        const tr = document.createElement("tr")
        addRowData(tr, workout.Date)
        addRowData(tr, `${workout.Distance}km`)
        addRowData(tr, workout.Duration)
        addRowData(tr, `${workout.Elevation}m`)
        addRowData(tr, workout.AvgPace)
        addRowData(tr, workout.AvgHR)
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
        if (length(workouts) == 0) {
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