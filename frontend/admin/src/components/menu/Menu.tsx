import './Menu.scss'
import {Link} from "react-router-dom";
import {MenuItems} from "./Preset.tsx";

const Menu = () => {
    return (
        <div className="menu">
            {MenuItems.map((item) => (
               <div className="item" key={item.id}>
                   <span className="title">
                       { item.title }
                   </span>
                   { item.listItems.map((listItem) => (
                       <Link to="/" className="listItem" key={listItem.id}>
                           <img src={listItem.icon} alt=""/>
                           <span className="listItemTitle">
                               { listItem.title }
                           </span>
                       </Link>
                   ))}
                </div>
            ))}
        </div>
    )
}

export default Menu
