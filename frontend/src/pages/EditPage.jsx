import React, { useEffect } from 'react'
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { Button, Form, notification, Input } from 'antd'
import { NavLink, useNavigate } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { useParams } from 'react-router-dom';
import { fetchAllAuctions, fetchAuctionPhotos, fetchBetsForAuction, makeBet, updateAuction  } from '../redux/features/auction/auctionSlice'


const { TextArea } = Input;

export const EditPage = () => {
    const { id } = useParams();
    const navigate = useNavigate();

    const dispatch = useDispatch()

  console.log("id is " + id);

  useEffect(() => {
    dispatch(fetchAllAuctions());
  }, [dispatch, id])


  const auction = useSelector(state =>
    state.auction.auctions.find(auction => auction.Id.toString() === id)
  );


  const onFinish = async (values) => {
    values.id = id
    try {
            console.log(values);

            dispatch(updateAuction(values))
        } catch (error) {
            console.log(error)
        }
        console.log('Received values of form: ', values);
        navigate(`/auction/${id}`)
        };
   
    const [form] = Form.useForm();
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

    <h1 className='w-full text-center' >Edit auction data</h1>

    <Form.Item
            name="title"
            label="Auction Title">
            <Input/>
        </Form.Item>

        <Form.Item 
            name="description"
            label="Description" >
            <TextArea rows={4}/>
        </Form.Item>

    <Form.Item>
      <Button type="primary" htmlType="submit" className="login-form-button w-full">
        Edit data
      </Button>
      Or <NavLink className="login-form-forgot" to={'/'}> return to listing</NavLink>
    </Form.Item>
  </Form>
  </div>
  )
}
