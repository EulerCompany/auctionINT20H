import React from 'react'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { Button, Form, Input } from 'antd'
import { NavLink } from 'react-router-dom'

export const RegistrationPage = () => {
  const onFinish = (values) => {
    console.log('Received values of form: ', values);
  };

  return (
    <div className='flex justify-center'>
    <Form
    name="login"
    className="login-form  w-3/12 pt-20"
    initialValues={{
      remember: true,
    }}
    onFinish={onFinish}>

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

    <Form.Item
      name="confirm_password"
      rules={[
        {
          required: true,
          message: 'Please confirm your Password!',
        },
      ]}>

      <Input
        prefix={<LockOutlined className="site-form-item-icon" />}
        type="password"
        placeholder="Confirm password"/>
    </Form.Item>

    <Form.Item>
      <Button type="primary" htmlType="submit" className="login-form-button w-full">
        Register me
      </Button>
      Allready have an acoount? <NavLink className="login-form-forgot" to={'/registration'}> Sign In now</NavLink>
    </Form.Item>
  </Form>
  </div>
  )
}
