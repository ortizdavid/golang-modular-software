package helpers

func RecoverLinkTmpl(userName string, recoverLink string) string {
	return `
		<html>
			<body>
				<h1>Password Recovery!</h1>
				<p>Hello, `+userName+`!</p>
				<p>To recover password Click <a href="`+recoverLink+`">Here</a></p>
			</body>
		</html>`
}

func RecoverPasswordTmpl(userName string, password string) string {
	return `
		<html>
			<body>
				<h1>Password Changed!</h1>
				<p>Hello, `+userName+`!</p>
				<p>Your new password: <b>`+password+`</b></p>
			</body>
		</html>`
}