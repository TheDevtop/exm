import axios from 'axios';

const addr = "http://exm.some.network:1800/search/object";

export class ResultForm {
    Error: string = "";
    Count: number = 0;
    Result: string[] = [];
}

export function Search(r: string, o: string): ResultForm {
    const bodyForm = {
        Object: o,
        Regex: r
    };

    axios.post(addr, bodyForm).then((resp) => {
        let rf = resp.data as ResultForm;
        return rf;
    }).catch((err) => {
        console.error(err);
    });

    return new ResultForm();
}