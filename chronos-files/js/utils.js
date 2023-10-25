function token() {

	const cookies = document.cookie.split(";").map(x => x.split('='))
	const token = cookies.find(x => x[0] == "jwt")[1]
	return token
}
