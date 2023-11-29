
let state = {
  schedules: []
}

function add_schedule(start, duration, week_day, day, original, type) {
  state.schedules.push({
    start: start,
    duration: duration,
    week_day: week_day,
	day: day,
	original: original,
    type: type 
  });
}

function sign_schedule(state) {
	console.log(state)

	const start = state.day + state.original.start.substring(10,22)
	const end = state.day + state.original.end.substring(10,22)
	req("/api/scheduling/add", {
		method: "POST",
		body: JSON.stringify({
			start: start,
			end: end,
			time_id: state.original.id,
			user_id: parseJwt(token()).user_id
		})
	})
}

function render_schedules() {
  let i = 1
  state.schedules.sort((st,e) => ((new Date(st.day)).getTime() + st.start)  - ((new Date(e.day)).getTime() + e.start) )
  for (let s of state.schedules) {
    const schedule_height = document.querySelector(".schedules").offsetHeight;

    const col = document.getElementById(s.week_day);
    const schedules = col.querySelector(".schedules");

    const week_index = ['seg', 'ter', 'qua', 'qui', 'sex'].indexOf(s.week_day);

    const el_top = s.start*schedule_height/11;
	

    const el = document.createElement("div") 
	el.tabIndex = i 
	el.style = `height: ${s.duration*(schedule_height/11)}px; transform: translate(0,${el_top}px)` 
	el.classList.add("schedule-item")
	if (s.type === "notfree")
		el.classList.add("notfree")
	if (s.type === "free")
		el.onclick = (e) => sign_schedule(s)

    schedules.appendChild(el)
  console.log(el)
    i++
  }
}

function handle_error(e) {
	console.log(e);
}

async function request_schedules(offset, forward=true) {
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

	const resfree = await req(`/api/time?date=${date}&page=0`, null, handle_error);
	const resnotfree = await req(`/api/scheduling?date=${date}&page=0`, null, handle_error);


	const freejson = await resfree.json();
	const notfreejson = await resnotfree.json();
	
	if (freejson.message)
		window.location = "/login.html";

	freejson.forEach(t => {
		const start = (new Date(t.start)).getHours() + (new Date(t.start)).getMinutes()/60 - 9;
		const end = (new Date(t.end)).getHours() + (new Date(t.end)).getMinutes()/60 - 9;
		const duration = end - start;
		add_schedule(start,duration, week_day_name,date,t, "free");
	});
	notfreejson.forEach(t => {
		const start = (new Date(t.start)).getHours() + (new Date(t.start)).getMinutes()/60 - 9;
		const end = (new Date(t.end)).getHours() + (new Date(t.end)).getMinutes()/60 - 9;
		const duration = end - start;
		add_schedule(start,duration, week_day_name,date,t, "notfree");
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
