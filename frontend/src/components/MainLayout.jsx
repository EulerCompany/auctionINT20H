import React from 'react'
import { NotificationOutlined } from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { isAuth, logout } from '../redux/features/auth/authSlice';
import { HeaderComponent }  from './HeaderComponent.jsx';
import { FooterComponent } from './FooterComponent.jsx';
import { SliderComponent } from './SliderComponent.jsx';

const { Header, Footer, Content, Sider } = Layout;


export const MainLayout = ({ children }) => {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <React.Fragment>
      <Layout>
        <HeaderComponent/>
      <Layout>
        <SliderComponent/>
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
        <FooterComponent/>
        </Layout>
      </Layout>
    </Layout>
    </React.Fragment>
  )
}
