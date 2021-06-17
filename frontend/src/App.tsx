import Header from "./components/Header";
import Home from "./pages/Home";
import Form from "./pages/Form";
import { BrowserRouter as Router, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";

var App: React.FC = () => {

    return (
        <>
            <Router>
                <Header />
                <Route exact path="/" component={Home}></Route>
                <Route exact path="/new" component={Form}></Route>
            </Router>
        </>
    )

}

export default App;