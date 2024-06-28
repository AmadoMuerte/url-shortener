import './UrlElement.css';

interface Props {
    id: number;
    alias: string;
    url: string;
    address: string;
}

function UrlElement({ id, alias, url }: Props) {
    return (
        <li className="urlElement" >
            <div title="Айди" className="urlElement__id">{id}</div>
            <div title="Алиас (нажмите что бы скопировать)" className="urlElement__alias">{alias}</div>
            <a title="Оригинальный URL" target="_blank" href={url} className="urlElement__url">{url}</a>
        </li>
    );
}

export default UrlElement;
