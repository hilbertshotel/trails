// COLORS
const TURQUOISE = "rgb(111, 218, 243)"
const BROWN = "#555759"
const LIGHT_GREY = "#f6f6f6"
const PINK = "#c85e7a"

// COLUMN FUNCTIONS
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

// SORT FUNCTIONS
const sortWorkouts = async (row, table, column, socket, direction) => {
    const marked = getMarked(row)
    if (column.id != marked.id) {
        unmarkColumn(marked)
        markColumn(column)
    }
    clearTable(table)
    socket.send(JSON.stringify({name: column.id, direction: direction}))
    const newDirection = direction == "desc" ? "asc" : "desc" 
    column.onclick = () => { sortWorkouts(row, table, column, socket, newDirection) }
}

// TABLE FUNCTIONS
const clearTable = (table) => {
    table.innerHTML = ""
} 

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

const loadSorters = (row, table, columns, socket) => {
    for (const name of columns) {
        const column = document.getElementById(name)
        column.onclick = () => { sortWorkouts(row, table, column, socket, "desc") }
    }
}

// ON PAGE LOAD
const main = async () => {
    
    // init constants
    const row = document.getElementById("first_row")
    const table = document.getElementById("workouts")
    const firstColumn = document.getElementById("date")
    const columns = ["date", "distance", "duration", "elevation", "pace", "hr"]

    // get ws server address
    const response = await fetch("/wsa")
    if (response.ok) {
        const addr = await response.text()

        // connect to ws server
        const socket = new WebSocket(addr)
        socket.onopen = (_) => {
            console.log("Web Socket Connection Established")
            socket.send(JSON.stringify({name: "date", direction: "desc"}))
        }
        socket.onclose = (_) => {
            console.log("Web Socket Connection Terminated")
        }

        // on receive
        socket.onmessage = (msg) => {
            const workouts = JSON.parse(msg.data)
            loadWorkouts(table, workouts)
        }

        // load sorters
        loadSorters(row, table, columns, socket)
        markColumn(firstColumn)
    }

    else {
        console.log(response.status, response.statusText)
    }
}

main()