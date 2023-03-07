import { Router, route } from 'preact-router'
import { render } from 'preact'
import { useEffect } from 'preact/hooks'
import { PageHomepage } from './pages.jsx'

const Redirect = ({ to }) => {
    useEffect(() => {
        route(to, true)
    }, [])

    return null
}

const Main = ({}) => {
    return (
        <Router>
            <PageHomepage path="/" />
            <Redirect default to="/" />
        </Router>
    )
}

render(<Main />, document.body)
