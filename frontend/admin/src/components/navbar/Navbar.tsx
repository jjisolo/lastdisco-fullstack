import './Navbar.scss'

const Navbar = () => {
    return (
        <div className="navbar">
           <div className="logo">
               <img src="logo.png" alt="" />
               <div className="nameMain">
                   <span className="sp-2">
                       Last
                       <span className="sp-1">
                           Disco
                       </span>
                       ™
                   </span>
               </div>
           </div>
           <div className="icons">
               <img src="/search.svg" alt="" className="icon" />
               <img src="/app.svg" alt="" className="icon" />
               <img src="/expand.svg" alt="" className="icon" />
               <div className="notification">
                   <img src="/notifications.svg" alt="" className="icon" />
                   <span>1</span>
               </div>
               <div className="user">
                   <img src="https://t3.ftcdn.net/jpg/02/43/12/34/360_F_243123463_zTooub557xEWABDLk0jJklDyLSGl2jrr.jpg" alt="" className="icon" />
                   <span>Issa</span>
               </div>
               <img src="/settings.svg" alt="" className="icon" />
           </div>
        </div>
    )
}

export default Navbar
