import axios from 'axios'
import React, { SyntheticEvent, useEffect, useState } from 'react'
import Wrapper from '../components/Wrapper'

export default function Profile() {

	const [first_name, setFirstName] = useState('')
	const [last_name, setLastName] = useState('')
	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const [password_confirm, setPasswordConfirm] = useState('')
	const [redirect, setRedirect] = useState(false)

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get('user')
                setFirstName(data.first_name)
                setLastName(data.last_name)
                setEmail(data.email)

            }
        )()
    }, [])

    const infoSubmit = async (e: SyntheticEvent) => {
        e.preventDefault()
        await axios.put('users/info',{
            first_name,
            last_name,
            email
        })
    }

    const passwordSubmit = async (e: SyntheticEvent) => {
        e.preventDefault()
        await axios.put('users/password',{
            password,
            password_confirm
        })
    }

    return (
        <Wrapper>
            <h3>Account Information</h3>
                <hr/>
                <form onSubmit={infoSubmit}>
                    <div className="form-group">
                        <label>First Name</label>
                        <input type="text" className="form-control" name="first_name"
                               defaultValue={first_name}
                               onChange={e => setFirstName(e.target.value)}
                        />
                    </div>
                    <div className="form-group">
                        <label>Last Name</label>
                        <input type="text" className="form-control" name="last_name"
                               defaultValue={last_name}
                               onChange={e => setLastName(e.target.value)}
                        />
                    </div>
                    <div className="form-group">
                        <label>Email</label>
                        <input type="text" className="form-control" name="email"
                               defaultValue={email}
                               onChange={e => setEmail(e.target.value)}
                        />
                    </div>

                    <button className="btn btn-outline-secondary">Save</button>
                </form>

                <h2 className="mt-4">Change Password</h2>
                <hr/>
                <form onSubmit={passwordSubmit}>
                    <div className="form-group">
                        <label>Password</label>
                        <input type="password" className="form-control" name="password"
                               onChange={e => setPassword(e.target.value)}
                        />
                    </div>
                    <div className="form-group">
                        <label>Password Confirm</label>
                        <input type="password" className="form-control" name="password_confirm"
                               onChange={e => setPasswordConfirm(e.target.value)}
                        />
                    </div>

                    <button className="btn btn-outline-secondary">Save</button>
                </form>
        </Wrapper>
    )
}
