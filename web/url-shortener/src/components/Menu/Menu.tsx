import './Menu.css';
import addIcon from '../../assets/add.svg';

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
            <li className="button menu__add" onClick={handleAddClick}>
                <img src={addIcon} alt="добавить"/>
            </li>
        </ul>
    );
}

export default Menu;