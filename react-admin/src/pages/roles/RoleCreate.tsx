import axios from 'axios'
import Wrapper from "../../components/Wrapper"
import React, { useState, useEffect, SyntheticEvent } from "react"
import { Permission } from "../../models/permission"
import { Redirect } from 'react-router'

export default function RoleCreate() {

    const [permissions, setPermissions] = useState([])
    const [selected, setSelected] = useState([] as number[])
    const [name, setName] = useState('')
    const [redirect, setRedirect] = useState(false)

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get('permissions')
                setPermissions(data)
            }
        )()
    },[])

    const check = (id: number) => {
        if (selected.some(s => s === id)) {
            setSelected(selected.filter(s => s !== id))
            return
        }
        setSelected([...selected, id])
    }

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault()
        await axios.post('roles',{
            name,
            permissions: selected

        })
        setRedirect(true)
        console.log(selected)
    }

    if (redirect) {
        return <Redirect to="/roles" />
    }

    return(
        <Wrapper>
            <form onSubmit={submit}>
                <div className="form-group row">
                    <label htmlFor="name" className="col-sm-2 col-form-label">Name</label>
                    <div className="col-sm-10">
                        <input type="text" className="form-control" name="name" id="name"
                            onChange={e => setName(e.target.value)}
                        />
                    </div>
                </div>

                <div className="form-group row">
                    <label className="col-sm-2 col-form-label">Permissions</label>
                    <div className="col-sm-10">
                        {permissions.map(
                            (p: Permission) => {
                                return (
                                    <div className="form-check form-check-inline col-3" key={p.id}>
                                        <input className="form-check-input" type="checkbox" value={p.id}
                                            onChange={() => check(p.id)}
                                        />
                                        <label className="form-check-label">{p.name}</label>
                                    </div>
                                )
                            }
                        )}
                    </div>
                </div>

                <button className="btn btn-outline-secondary">Save</button>
            </form>
        </Wrapper>
    )
}