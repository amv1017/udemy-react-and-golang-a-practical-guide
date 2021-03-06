import React, { SyntheticEvent, useState } from 'react'
import axios from 'axios'
import { Redirect } from 'react-router-dom'

export default function Login() {

	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [redirect, setRedirect] = useState(false)

	const submit = async (e: SyntheticEvent) => {

		e.preventDefault()
		const { data } = await axios.post('login', {
			email,
			password
		})

		setRedirect(true)

	}

	if (redirect) {
		return <Redirect to={'/'} />
	}

	return (
		<main className="form-signin">
			<form onSubmit={submit}>
				<h1 className="h3 mb-3 fw-normal">Please sign in</h1>

				<input type="email" className="form-control" placeholder="user@email.com" required
					onChange={e => setEmail(e.target.value)}
				/>

				<input type="password" className="form-control" placeholder="password" required
					onChange={e => setPassword(e.target.value)}
				/>

				<button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
			</form>
	</main>
	)
}
