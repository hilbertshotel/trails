// COLORS
const TURQUOISE = "#6fdaf3"
const BROWN = "#555759"
const LIGHT_GREY = "#f6f6f6"
const PINK = "#c85e7a"

// COLUMN FUNCTIONS
const getMarked = () => {
    const firstRow = document.getElementById("first_row")
    for (const column of firstRow.children) {
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

// SORT FUNCTIONS
const sortWorkouts = (table, workouts, column) => {
    const marked = getMarked()
    if (column.id == marked.id) {
        // HOW TO HANDLE REVERSE SORT ON THE SAME COLUMN ? HOW TO KEEP STATE ?
    }

    else {
        sortByKey(workouts, column.id)
        unmarkColumn(marked)
        clearTable(table)
        loadWorkouts(table, workouts)
        markColumn(column)
    }
}

const sortByKey = (objectArray, key) => {

}

// TABLE FUNCTIONS
const clearTable = (table) => {
    table.innerHTML = ""
} 

const addRowData = (row, data) => {
    td = document.createElement("td")
    td.innerHTML = data
    row.append(td)
}

const loadWorkouts = (table, workouts) => {
    for (const workout of workouts) {
        const row = document.createElement("tr")

        addRowData(row, workout.Date)
        addRowData(row, `${workout.Distance}km`)
        addRowData(row, workout.Duration)
        addRowData(row, `${workout.Elevation}m`)
        addRowData(row, workout.AvgPace)
        addRowData(row, workout.AvgHR)

        table.append(row)
    }
}

const loadSorters = (table, columns, workouts) => {
    for (const name of columns) {
        const column = document.getElementById(name)
        column.onclick = () => { sortWorkouts(table, workouts, column) }
    }
}

// ON PAGE LOAD
const main = async () => {
    const table = document.getElementById("workouts")
    const date = document.getElementById("date")
    const columns = ["date", "distance", "duration", "elevation", "pace", "hr"]

    let response = await fetch("/workouts")
    if (response.ok) {
        let workouts = await response.json()

        // load workouts
        loadWorkouts(table, workouts.reverse())

        // load sorters
        loadSorters(table, columns, workouts)

        // mark first column
        markColumn(date)
    }
    
    else {
        console.log(response.statusText)
    }
}

main()