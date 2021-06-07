import React, { useState } from 'react'

export default function Login() {

    const [count, setCount] = useState(0)

    return (
        <div className="container">
            <h1>Count {count}</h1>
            <input type="number" onChange={e => setCount(parseInt(e.target.value))}></input>
        </div>
    )
}
