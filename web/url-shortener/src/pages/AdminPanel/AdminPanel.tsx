import UrlList from "../../components/UrlList/UrlList.tsx";
import Menu from "../../components/Menu/Menu.tsx";
import PopupForm from "../../components/PopupForm/PopupForm.tsx";
import {useEffect, useState} from "react";
import {useCookies} from "react-cookie";



interface UrlResponse {
    status: string
    data: UrlInfo[]
    address: string
}

interface UrlInfo {
    id: number
    alias: string
    url: string
}

async function getAliases(login:  string, password: string) {
    const response = await fetch("http://localhost:8082/url/all", {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(`${login}:${password}`)
        },
    })
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}


function AdminPanel() {
    const [popupIsOpen, setPopupIsOpen] = useState<boolean>(false);
    const [cookies] = useCookies(['login', 'password']);
    const [data, setData] = useState<UrlResponse>();
    const [isUpdating, setUpdating] = useState<boolean>(false);

    useEffect(() => {
        if (!isUpdating) {
            const fetchAliases = async () => {
                const data = await getAliases(cookies.login, cookies.password);
                setData(data);
                console.log(data)
            };
            fetchAliases();
        }

    }, [cookies.login, cookies.password, isUpdating]);


    return (
        <div className="container">
            <Menu setPopupIsOpen={setPopupIsOpen}/>
            <UrlList data={data}/>
            <PopupForm popupIsOpen={popupIsOpen} setPopupIsOpen={setPopupIsOpen} setUpdating={setUpdating} />
        </div>
    );
}

export default AdminPanel;