
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

add_schedule(2, 2, "seg", "free")
add_schedule(0, 3, "ter", "free")
add_schedule(0, 4, "qua", "free")
add_schedule(0, 6, "qui", "free")
add_schedule(0, 1, "sex", "free")
render_schedules()
