import React, { useEffect } from 'react'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { Button, Form, Input, notification } from 'antd'
import { NavLink, useNavigate } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { isAuth, registerUser } from '../redux/features/auth/authSlice'

export const RegistrationPage = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  // const checkIsAuth = useSelector(isAuth)
  const checkIsAuth = false



  const { status } = useSelector((state) => state.auth)

  const onFinish = (values) => {
    try {
      // dispatch(registerUser(values))
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
  }, [status, checkIsAuth, navigate])

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
   
    onFinish={onFinish}>

    <h1 className='w-full text-center' >Registration</h1>

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

      <Input prefix={<LockOutlined className="site-form-item-icon" />} type="password" placeholder="Password"/>
    </Form.Item>

    <Form.Item>
      <Button type="primary" htmlType="submit" className="login-form-button w-full">
        Register me
      </Button>
      Allready have an acoount? <NavLink className="login-form-forgot" to={'/login'}> Sign In now</NavLink>
    </Form.Item>
  </Form>
  </div>
  )
}
