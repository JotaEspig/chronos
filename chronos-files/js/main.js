
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
    const hours_width = document.querySelector(".time").offsetWidth
    const schedule_width = document.querySelector(".schedule").offsetWidth
    const schedule_height = document.querySelector(".schedules").offsetHeight

    const col = document.getElementById(s.week_day)
    const schedules = col.querySelector(".schedules")

    const week_index = ['seg', 'ter', 'qua', 'qui', 'sex'].indexOf(s.week_day) 
    console.log(week_index)
    console.log("test")

    const el_width = 100
    const el_offset = hours_width + ((week_index + week_index+1)*(schedule_width/5))/2 - el_width/2
    const el = `
                    <div style="left: ${el_offset}px; height: ${s.duration*(schedule_height/11)}px" class="schedule-item">

                    </div>
                    `

    schedules.innerHTML += el
  }
}

add_schedule(0, 2, "seg", "free")
add_schedule(0, 3, "ter", "free")
add_schedule(0, 4, "qua", "free")
add_schedule(0, 6, "qui", "free")
add_schedule(0, 1, "sex", "free")
render_schedules()
