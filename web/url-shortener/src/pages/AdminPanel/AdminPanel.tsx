import UrlList from "../../components/UrlList/UrlList.tsx";
import Menu from "../../components/Menu/Menu.tsx";
import PopupForm from "../../components/PopupForm/PopupForm.tsx";
import {useEffect, useState} from "react";
import {useCookies} from "react-cookie";
import {PopUpConfig, PopUpData, UrlResponse} from "../../types.ts";



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
    const [data, setData] = useState<UrlResponse>({data: [], status: "", address: ""});
    const [isUpdating, setUpdating] = useState<boolean>(false);
    const [popUpData, setPopUpData] = useState<PopUpData>({url:  "", alias: ""});
    const [editId, setEditId] = useState<number>(0);

    useEffect(() => {
        if (!isUpdating) {
            const fetchAliases = async () => {
                const data = await getAliases(cookies.login, cookies.password);
                if (data.status === "OK") {
                    setData(data);
                    setUpdating(true);
                }
            };
            fetchAliases()
        }

    }, [cookies.login, cookies.password, isUpdating]);

    const popUpConfig: PopUpConfig = {
        popupIsOpen: popupIsOpen,
        setPopupIsOpen,
        setUpdating,
        setPopUpData,
        popUpData: popUpData,
        editId,
        setEditId: setEditId
    }

    return (
        <div className="container">
            <Menu setPopupIsOpen={setPopupIsOpen}/>
            <UrlList
                data={data}
                popUpConfig={popUpConfig}
            />
            <PopupForm
                popUpConfig={popUpConfig}
            />
        </div>
    );
}

export default AdminPanel;