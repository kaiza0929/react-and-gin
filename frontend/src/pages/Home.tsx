import { useEffect } from "react";
import axios from "axios";

var Home: React.FC = () => {

    useEffect(() => {
        axios.get("/api/log", {headers: {"Content-type": "application/json"}})
        .then((res) => alert(JSON.stringify(res)))
        .catch((err) => alert(err))
    }, []);

    return (
        <p>this is home page</p>
    )

}

export default Home;