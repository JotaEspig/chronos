function login(e) {
    e.preventDefault();
    console.log(e);
    const username = document.querySelector("#username").value;
    const password = document.querySelector("#password").value;

    const api = "/api/login";
	req("/api/login", { 
		method: "POST", 
        body: JSON.stringify({
            username: username,
            password: password
        })
	}).then(res => res.json())
        .then(res => {
            if (res.message) {
                alert("failed: " + res.message);
                return;
            }
            const token = res.token;
            document.cookie =`jwt=${token}`;
            window.location = "/";
        });
}
