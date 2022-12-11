// ON PAGE LOAD
const loadWorkouts = (workouts) => {
    const table = document.getElementById("workouts")
    
    for (const workout of workouts) {
        const tr = document.createElement("tr")
        // date
        let td = document.createElement("td")
        td.innerHTML = workout.Date
        tr.append(td)
        // distance
        td = document.createElement("td")
        td.innerHTML = workout.Distance
        tr.append(td)
        // duration
        td = document.createElement("td")
        td.innerHTML = workout.Duration
        tr.append(td)
        // elevation
        td = document.createElement("td")
        td.innerHTML = workout.Elevation
        tr.append(td)
        // avg pace
        td = document.createElement("td")
        td.innerHTML = workout.AvgPace
        tr.append(td)
        // avg hr
        td = document.createElement("td")
        td.innerHTML = workout.AvgHR
        tr.append(td)

        table.append(tr)
    }
}

const main = async () => {
    let response = await fetch("/workouts")
    if (response.ok) {
        let workouts = await response.json()
        loadWorkouts(workouts.reverse())
    }
}

main()