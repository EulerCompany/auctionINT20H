
import React from 'react'
import { NotificationOutlined } from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { isAuth, logout } from '../redux/features/auth/authSlice';
import { HeaderComponent }  from './HeaderComponent.jsx';
import { FooterComponent } from './FooterComponent.jsx';

const { Header, Footer, Content, Sider } = Layout;



export function SliderComponent() {
  const checkIsAuth = useSelector(isAuth)
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

    const userId = useSelector((state) => {
        return state.auth.user;
    });
    if (!checkIsAuth) {
        return null
    }
    

  return (
      <Sider
          collapsible
          width={200}
          style={{
            background: colorBgContainer,
          }}>

          <Menu
            mode="inline"
            defaultSelectedKeys={[window.location.pathname]}
            selectedKeys={[window.location.pathname]}
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
                    key: '/user/:id/auctions',
                    label: <NavLink to={`/user/${userId}/auctions`} >My auctions</NavLink>,
              }],
          },
          ]}/>
        </Sider>
  )
}
