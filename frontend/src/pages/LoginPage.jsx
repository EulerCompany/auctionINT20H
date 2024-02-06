import React, { useEffect } from 'react'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { Button, Form, notification, Input } from 'antd'
import { NavLink, useNavigate } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { isAuth, loginUser } from '../redux/features/auth/authSlice'


export const LoginPage = () => {

  const dispatch = useDispatch()
  const navigate = useNavigate()
  // const checkIsAuth = useSelector(isAuth)
  const checkIsAuth = true

  const { status } = useSelector((state) => state.auth)

  const onFinish = (values) => {
    try {
      // dispatch(loginUser(values))
    } catch (error) {
      console.log(error)
    }
    console.log('Received values of form: ', values);
    form.resetFields()
  };

  useEffect(() => {
    if (status) {
      openNotification('topRight')
    }
    if (checkIsAuth) {
      navigate('/')
    }
  }, [status, checkIsAuth, navigate, openNotification])

  const [api] = notification.useNotification();
  const [form] = Form.useForm();

  const openNotification = (placement) => {
    api.info({
      message: `Wow!`,
      description: status,
      placement,
    })};

  return (
    <div className='flex justify-center'>
    <Form
    form={form}
    name="login"
    className="login-form  w-3/12 pt-20"
    initialValues={{
      remember: true,
    }}
    onFinish={onFinish}>

    <h1 className='w-full text-center' >Sign In</h1>

    <Form.Item
      name="username"
      rules={[
        {
          required: true,
          message: 'Please input your Username!',
        },
      ]}>

      <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="Username" />
    </Form.Item>

    <Form.Item
      name="password"
      rules={[
        {
          required: true,
          message: 'Please input your Password!',
        },
      ]}>

      <Input
        prefix={<LockOutlined className="site-form-item-icon" />}
        type="password"
        placeholder="Password"/>
    </Form.Item>

    <Form.Item>
      <Button type="primary" htmlType="submit" className="login-form-button w-full">
        Sign In
      </Button>
      Or <NavLink className="login-form-forgot" to={'/registration'}> register now</NavLink>
    </Form.Item>
  </Form>
  </div>
  )
}
