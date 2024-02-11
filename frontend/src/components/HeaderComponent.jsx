
import React from 'react'
import { Button, Layout, Menu } from 'antd';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { isAuth, logout } from '../redux/features/auth/authSlice';

const { Header } = Layout;

export function HeaderComponent() {
  const checkIsAuth = useSelector(isAuth)

  const dispatch = useDispatch()

  const location = useLocation();
  const navigate = useNavigate();

  const logoutHandler = () => {
    dispatch(logout())
    window.localStorage.removeItem('token')
    navigate('/')
  }
  const createAuctionHandler = () => {
    navigate('/create')
  }

  return (
      <Header
        style={{
          display: 'flex',
          justifyContent: 'space-between' ,
          alignItems: 'center',
        }}>

        <div className="text-3xl"> 
        <a className='text-white' href='/'> AUCTION 423 </a>
        </div>

        {!checkIsAuth && <Menu
        className='w-2/3 flex justify-end'
          theme="dark"
          mode="horizontal"
          selectedKeys = {[location.pathname]}
          defaultSelectedKeys={location.pathname}

          items={[
          {
            key: '/login',
            label: <NavLink to={'/login'} >Sign In</NavLink>,
          },
          {
            key: '/registration',
            label: <NavLink to={'/registration'} >Sign Up</NavLink>,
          },
          ]}

          style={{
            alignItems: 'center',
            minWidth: 0,
          }}/> }
         {checkIsAuth && <Button className='mr-2' type="primary" onClick={createAuctionHandler} >Create Auction</Button>}
         {checkIsAuth && <Button type="primary" onClick={logoutHandler} >Log out</Button>} 
      </Header>
    )
}

