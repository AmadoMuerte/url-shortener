import './UrlElement.css';
import {PopUpConfig, URLInfo} from "../../types.ts";
import editIcon from "../../assets/edit.svg";
import deleteIcon from "../../assets/delete.svg"
import {deleteUrl} from "../../lib/api/url.ts";
import {useCookies} from "react-cookie";

type Props = {
    urlInfo: URLInfo
    config : Config
}

type Config =  {
    popUpConfig:  PopUpConfig,
}

function UrlElement(props: Props) {
    const { urlInfo, config} = props
    const [cookie,] = useCookies(['login', 'password']);

    const handleEditClick = () => {
        config.popUpConfig.setPopupIsOpen(true)
        config.popUpConfig.setPopUpData({
            url: urlInfo.url,
            alias: urlInfo.alias
        })
        config.popUpConfig.setEditId(urlInfo.id)
    }

    const handleCopy = () => {
        navigator.clipboard.writeText("http://" + urlInfo.address + '/' + urlInfo.alias)
            .then(() => {
                alert('Ссылка скопирована.')
            })
    }

    const handleDeleteClick = () => {
        config.popUpConfig.setUpdating(true)
        deleteUrl(cookie.login, cookie.password, urlInfo.id)
            .then((res) => {
                if (res.status === "OK") {
                    config.popUpConfig.setUpdating(false)
                } else {
                    alert(res.error)
                }
            })
    }

    return (
        <li className="urlElement" >
            <div title="Айди" className="urlElement__id">{urlInfo.id}</div>
            <div
                title="Алиас (нажмите что бы скопировать)"
                className="urlElement__alias"
                onClick={handleCopy}
            >{urlInfo.alias}</div>
            <a title="Оригинальный URL" target="_blank" href={urlInfo.url} className="urlElement__url">{urlInfo.url}</a>
            <button className="button urlElement__delete" onClick={handleDeleteClick}>
                <img src={deleteIcon} alt="удалить"/>
            </button>
            <button className="button urlElement__edit" onClick={handleEditClick}>
                <img src={editIcon} alt={'редактировать'}/>
            </button>
        </li>
    );
}

export default UrlElement;
