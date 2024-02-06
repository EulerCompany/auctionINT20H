import React from 'react'
import { NotificationOutlined } from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { NavLink, useLocation} from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { isAuth, logout } from '../redux/features/auth/authSlice';

const { Header, Footer, Content, Sider } = Layout;


export const MainLayout = ({ children }) => {

  // const checkIsAuth = useSelector(isAuth)
  const checkIsAuth = false

  const dispatch = useDispatch()

  const location = useLocation();

  const logoutHandler = () => {
    dispatch(logout())
    window.localStorage.removeItem('token')
  }

  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <React.Fragment>
      <Layout>
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
         {checkIsAuth && <Button type="primary" onClick={logoutHandler} >Log out</Button> }
      </Header>

      <Layout>
      {checkIsAuth && <Sider
          width={200}
          style={{
            background: colorBgContainer,
          }}>

          <Menu
            mode="inline"
            defaultSelectedKeys={location.pathname}
            selectedKeys={[location.pathname]}
            defaultOpenKeys={['auctions']}
            style={{
              height: '100%',
              borderRight: 0,
            }}
            items={[{
              key: 'auctions',
              icon: React.createElement(NotificationOutlined),
              label: 'Auctions',
              children: 
              [{
                  key: '/',
                  label: <NavLink to={'/'} >All auctions</NavLink>,
                },
                {
                  key: '/my-auctions',
                  label: <NavLink to={'/my-auctions'} >My auctions</NavLink>,
              }],
          },
          ]}/>
        </Sider>}

        <Layout
          style={{
            padding: '0 24px 24px',
          }}>

          <Content
            style={{
              padding: 24,
              margin: 0,
              minHeight: 780,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}>

            { children }

          </Content>

          <Footer style={{ textAlign: 'center' }}>
          int20h ©{new Date().getFullYear()} Created by Vlad and Maks
        </Footer>
        </Layout>
      </Layout>
    </Layout>
    </React.Fragment>
  )
}
