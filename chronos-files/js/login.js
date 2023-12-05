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
                document.getElementById("form").reset();
                document.getElementById("username").focus();

                Swal.fire({
                    position: "top-end",
                    title: "Erro",
                    icon: "error",
                    text: "O usuário e/ou a senha estão incorretos",
                    width: 400,
                    showConfirmButton: false,
                    timer: 1500
                });
                setInterval(() => {
                    document.getElementById("error").innerHTML = "O usuário e/ou a senha estão incorretos";
                }, 1500);
                return;
            }
            const token = res.token;
            document.cookie =`jwt=${token}`;
            window.location = "/";
        });
}
