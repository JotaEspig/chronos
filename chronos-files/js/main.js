
let state = {
  schedules: []
}

function add_schedule(start, duration, week_day, type) {
  state.schedules.push({
    start: start,
    duration: duration,
    week_day: week_day,
    type: type 
  });
}

function sign_schedule(state) {
	console.log(state)
}

function render_schedules() {
  for (let s of state.schedules) {
    const schedule_height = document.querySelector(".schedules").offsetHeight;

    const col = document.getElementById(s.week_day);
    const schedules = col.querySelector(".schedules");

    const week_index = ['seg', 'ter', 'qua', 'qui', 'sex'].indexOf(s.week_day);

    const el_top = s.start*schedule_height/11;
	

    const el = document.createElement("div") 
	el.style = `height: ${s.duration*(schedule_height/11)}px; transform: translate(0,${el_top}px)` 
	el.classList.add("schedule-item")
	el.onclick = (e) => sign_schedule(s)

    schedules.appendChild(el)
  }
}

function handle_error(e) {
	console.log(e);
}

async function request_schedules(offset, forward=true) {
	const api = "http://localhost:8080/api/time";
	const today = new Date((new Date()).setDate((new Date().getDate()) + offset));
	console.log(today);
	const year = today.getFullYear().toString();
	const month = (today.getMonth() + 1).toLocaleString('en-US', {
    minimumIntegerDigits: 2,
    useGrouping: false });


	const day = today.getDate().toLocaleString('en-US', {
    minimumIntegerDigits: 2,
    useGrouping: false });

	const week_day = today.getDay();

	const week_day_name = ['dom', 'seg', 'ter', 'qua', 'qui', 'sex', "sab"][week_day];

	const date = `${year}-${month}-${day}`;
	console.log(date);

	const req = await fetch(`${api}?date=${date}&page=0`, {
		headers: {
			"Authorization": "Bearer " + token()
		}
	}).catch(handle_error);


	const json = await req.json();
	console.log(json);
	
	if (json.message)
		window.location = "/login.html";

	json.forEach(t => {
		const start = (new Date(t.start)).getHours() + (new Date(t.start)).getMinutes()/60 - 9;
		const end = (new Date(t.end)).getHours() + (new Date(t.end)).getMinutes()/60 - 9;
		const duration = end - start;
		add_schedule(start,duration, week_day_name, "free");
	});

	if (week_day_name !=="sex" && forward)
		await request_schedules(offset + 1, forward);
	if (week_day_name !== "seg" && !forward)
		await request_schedules(offset - 1, forward);
}

async function change_week(el, n) {
	document.querySelectorAll(".week-selection button").forEach(el => el.classList.remove("current-week"));
	el.classList.add("current-week");

	document.querySelectorAll(".schedule-item").forEach(e => e.remove());

	state.schedules = [];

	await request_schedules(7*n, true);
	await request_schedules(7*n, false);
	render_schedules();

}


(async function() {

await request_schedules(0, true);
await request_schedules(0, false);
console.log(state);

render_schedules();
})();
