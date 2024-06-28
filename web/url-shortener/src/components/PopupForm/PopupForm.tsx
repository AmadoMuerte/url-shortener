import styles from './PopupForm.module.css'
import React, {useState} from "react";
import {useCookies} from "react-cookie";

interface Props {
    popupIsOpen: boolean
    setPopupIsOpen: (value: (((prevState: boolean) => boolean) | boolean)) => void,
    setUpdating: (value: (((prevState: boolean) => boolean) | boolean)) => void
}

interface Event {
    target: { value: React.SetStateAction<string>; };
}

async function addUrl(login:  string, password: string, alias: string, url: string) {
    const response = await fetch("http://localhost:8082/url", {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Basic ' + btoa(`${login}:${password}`)
        },
        body: JSON.stringify({
            alias: alias,
            url: url
        }),
    })
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return await response.json();
}

function PopupForm({popupIsOpen, setPopupIsOpen, setUpdating}: Props) {
    const [cookies] = useCookies(['login', 'password']);
    const [alias, setAlias] = useState("");
    const [url, setUrl] = useState("");

    if (!popupIsOpen) return null;

    const closePopup = (e: React.MouseEvent) => {
        if (e.target === e.currentTarget) {
            setPopupIsOpen(false);
        }
    };

    const handleAddClick = () => {
        setUpdating(true);
        addUrl(cookies.login, cookies.password, alias, url).then((response) =>
        {
            if (response.status === "OK") {
                setUpdating(false);
                setPopupIsOpen(false);
                clearFields();
            } else {
                alert(response.error);
            }
        })
    }

    const handleSetAlias = (event: Event) => {
        setAlias(event.target.value);
    }

    const handleSetUrl = (event: Event) => {
        setUrl(event.target.value);
    }

    const clearFields = () => {
        setAlias("");
        setUrl("");
    }

    return (
        <div  onClick={closePopup} className={styles.popupOverlay}>
            <div className={styles.popupContent}>
                <input title="Алиас" type="text" placeholder="Желаемый алиас"  value={alias} onChange={handleSetAlias}/>
                <input title="Ссылка" type="text" placeholder="Ссылка"  value={url}  onChange={handleSetUrl}/>
                <button onClick={handleAddClick}>Добавить</button>
            </div>
        </div>
    );
}

export default PopupForm;