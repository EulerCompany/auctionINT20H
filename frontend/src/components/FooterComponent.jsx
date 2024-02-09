
import React from 'react'
import { NotificationOutlined } from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { isAuth, logout } from '../redux/features/auth/authSlice';
import { HeaderComponent }  from './HeaderComponent.jsx';

const { Header, Footer, Content, Sider } = Layout;

export function FooterComponent() {
    return (
        <Footer style={{ textAlign: 'center' }}>
            int20h ©{new Date().getFullYear()} Created by Vlad and Maks
        </Footer>
    )
}
