
let state = {
  schedules: []
}

function add_schedule(start, duration, week_day, type) {
  state.schedules.push({
    start: start,
    duration: duration,
    week_day: week_day,
    type: type 
  })
}

function render_schedules() {
  for (let s of state.schedules) {
    const schedule_height = document.querySelector(".schedules").offsetHeight

    const col = document.getElementById(s.week_day)
    const schedules = col.querySelector(".schedules")

    const week_index = ['seg', 'ter', 'qua', 'qui', 'sex'].indexOf(s.week_day) 
    console.log(week_index)
    console.log("test")

    const el_top = s.start*schedule_height/11
    const el = `
                    <div style="height: ${s.duration*(schedule_height/11)}px; transform: translate(0,${el_top}px)" class="schedule-item">

                    </div>
                    `

    schedules.innerHTML += el
  }
}

async function request_schedules(week) {
	const api = "http://localhost:8080/api/time"
	const today = new Date(Date.now())
	const year = today.getFullYear().toString()
	const month = (today.getMonth() + 1).toLocaleString('en-US', {
    minimumIntegerDigits: 2,
    useGrouping: false })

	const week_inc = ['seg', 'ter', 'qua', 'qui', 'sex'].indexOf(week) + 1

	const day = (today.getDay() + week_inc).toLocaleString('en-US', {
    minimumIntegerDigits: 2,
    useGrouping: false })

	const date = `${year}-${month}-${day}`
	console.log(date)

	const req = await fetch(`${api}?date=${date}&page=0`)

	const json = await req.json()
	console.log(json)
	
	json.forEach(t => {
		const start = (new Date(t.start)).getHours() - 9
		const end = (new Date(t.end)).getHours() - 9
		const duration = end - start
		add_schedule(start,duration, week, "free")
	})

}

(async function() {

await request_schedules("seg")
await request_schedules("ter")
await request_schedules("qua")
await request_schedules("qui")
await request_schedules("sex")
console.log(state)

render_schedules()
})()
