import { useState } from "react";
import axios from "axios";

var Form: React.FC = () => {

    var [state, setState] = useState<{tag: string, content: string, result: string}>({
        tag: "",
        content: "",
        result: ""
    });

    return (
        <div>
            <label>タグ</label><input type="text" onKeyUp={(e: React.KeyboardEvent<HTMLInputElement>) => setState({...state, tag: e.currentTarget.value})}/>
            <textarea placeholder="" onKeyUp={(e: React.KeyboardEvent<HTMLTextAreaElement>) => setState({...state, content: e.currentTarget.value})}></textarea>
            <textarea placeholder="" onKeyUp={(e: React.KeyboardEvent<HTMLTextAreaElement>) => setState({...state, result: e.currentTarget.value})}></textarea>
            <input type="button" className="btn btn-primary" value="送信" onClick={() => {
               axios.post("http://localhost:8000/api/log", state, {headers: {"Content-Type": "application/json"}})
               .then((res) => alert(JSON.stringify(res)))
               .catch((err) => alert(err));
            }} /> 
        </div>
    )

}

export default Form;