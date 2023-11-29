function token() {
	const cookies = document.cookie.split(";").map(x => x.split('='));
	const arr = cookies.find(x => x[0] == "jwt");
    const token = arr ? arr[1] : "";
	return token;
}

function redirect(location) {
    window.location = location;
}


async function req(route, opts, handle_error) {
	const api = "" + route;

	const req = await fetch(api, {
		headers: {
			"Authorization": "Bearer " + token()
		},
		...opts
	}).catch(handle_error);

	return req
}
