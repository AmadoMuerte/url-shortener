import './UrlList.css'
import UrlElement from "../UrlElement/UrlElement.tsx";

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


function UrlList({data}: {data?: UrlResponse}) {

    return (
        <>
        <h2>Список ссылок</h2>
        <ul className="tableContainer">
            {data?.data.map((element: UrlInfo) => (
                <UrlElement id={element.id} alias={element.alias} url={element.url}  address={data.address} />
            ))}
        </ul>
        </>
    );
}

export default UrlList;