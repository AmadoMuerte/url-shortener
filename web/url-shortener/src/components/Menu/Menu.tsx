import './Menu.css';

interface Props {
    setPopupIsOpen?: (value: (((prevState: boolean) => boolean) | boolean)) => void
}

function Menu({setPopupIsOpen}: Props) {

    function handleAddClick() {
        if (setPopupIsOpen) {
            setPopupIsOpen(true)
        }
    }

    return (
        <ul className="menu">
            <li onClick={handleAddClick}>Add</li>
        </ul>
    );
}

export default Menu;