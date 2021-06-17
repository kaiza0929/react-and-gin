import { Link } from "react-router-dom";

var Header = () => {
    return (
        <div>
            <h1>this is header</h1>
            <Link to="/">ホーム</Link>
            <Link to="/new">テストログを追加</Link>
        </div>
    )
}

export default Header