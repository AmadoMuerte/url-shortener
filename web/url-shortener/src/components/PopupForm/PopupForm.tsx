import styles from './PopupForm.module.css'
import {useCookies} from "react-cookie";
import {addUrl, editUrl} from "../../lib/api/url.ts";
import {PopUpConfig} from "../../types.ts";
import React from "react";

function PopupForm({popUpConfig}: {popUpConfig: PopUpConfig}) {

    const [cookies] = useCookies(['login', 'password']);
    const popUpData = popUpConfig.popUpData;
    const setPopUpData = popUpConfig.setPopUpData;

    if (!popUpConfig.popupIsOpen) return null;


    const closePopup = (e: React.MouseEvent) => {
        if (e.target === e.currentTarget) {
            popUpConfig.setPopupIsOpen(false);
            clearFields();
        }
    };

    const handleAddClick = () => {
        popUpConfig.setUpdating(true);

        addUrl(cookies.login, cookies.password, popUpData).then((response) =>
        {
            if (response.status === "OK") {
                clearFields();
                popUpConfig.setUpdating(false);
                popUpConfig.setPopupIsOpen(false);
            } else {
                alert(response.error);
            }
        })
    }

    const handleEditClick = (id: number) => {
        popUpConfig.setUpdating(true);
        editUrl(cookies.login, cookies.password, popUpData, id).then((response) =>
        {
            if (response.status === "OK") {
                clearFields();
                popUpConfig.setEditId(0);
                popUpConfig.setUpdating(false);
                popUpConfig.setPopupIsOpen(false);
            } else {
                alert(response.error);
            }
        })
    }

    const handleSetAlias = (alias: string) => {
        setPopUpData({...popUpData, alias: alias});
    }

    const handleSetUrl = (url: string) => {
        setPopUpData({...popUpData, url: url});
    }

    const clearFields = () => {
        setPopUpData({alias: "", url: ""})
    }

    return (
        <div  onClick={closePopup} className={styles.popupOverlay}>
            <div
                className={styles.popupContent}
                onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                        popUpConfig.editId ? handleEditClick(popUpConfig.editId) : handleAddClick()
                    }
                }}
            >
                <input
                    title="Алиас"
                    type="text"
                    placeholder="Желаемый алиас"
                    value={popUpData.alias}
                    onChange={e => handleSetAlias(e.target.value)}/>
                <input
                    title="Ссылка"
                    type="text"
                    placeholder="Ссылка"
                    value={popUpData.url}
                    onChange={e => handleSetUrl(e.target.value)}
                />
                <button
                    onClick={() => popUpConfig.editId ? handleEditClick(popUpConfig.editId) : handleAddClick()}>
                    Добавить
                </button>
            </div>
        </div>
    );
}

export default PopupForm;