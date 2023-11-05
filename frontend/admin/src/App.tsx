import Home from './pages/home/Home.tsx'

import {
  createBrowserRouter,
  RouterProvider,
  Outlet,
} from "react-router-dom";
import Users from "./pages/users/Users.tsx";
import Products from "./pages/products/Products.tsx";
import Navbar from "./components/navbar/Navbar.tsx";
import Footer from "./components/footer/Footer.tsx";
import Menu from "./components/menu/Menu.tsx";
import './styles/global.scss'

function App() {

  const Layout = () => {
    return (
        <div className="main">
          <Navbar />
          <div className="container">
            <div className="menuContainer">
              <Menu/>
            </div>
            <div className="contentContainer">
              <Outlet />
            </div>
          </div>
          <Footer />
        </div>
    )
  }

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Layout/>,
      children: [
        {
          path: "/",
          element: <Home/>
        },
        {
          path: "/user-management",
          element: <Users />
        },
        {
          path: "/product-management",
          element: <Products />
        },
      ]
    }
  ]);

  return (
      <RouterProvider router={router} />
  )
}

export default App
