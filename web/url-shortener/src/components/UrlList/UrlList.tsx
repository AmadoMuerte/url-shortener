import './UrlList.css'
import UrlElement from "../UrlElement/UrlElement.tsx";
import {PopUpConfig, UrlResponse} from "../../types.ts";

interface UrlInfo {
    id: number
    alias: string
    url: string
}

function UrlList(props: {data?: UrlResponse, popUpConfig: PopUpConfig}) {
    const {data, popUpConfig} = props


    if (data?.data) {
        return (
            <>
                <h2>Список ссылок</h2>
                <ul className="tableContainer">
                    {data?.data.map((element: UrlInfo) => (
                        <UrlElement
                            key={element.id}
                            urlInfo={{
                                id: element.id,
                                alias: element.alias,
                                url: element.url,
                                address: data?.address
                            }}
                            config={{
                                popUpConfig: popUpConfig
                            }}
                        />
                    ))}
                </ul>
            </>
        );
    }

    return (
        <>
            <h2>Список пуст</h2>
        </>
    );

}

export default UrlList;