
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { NavLink, useNavigate } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { isAuth } from '../redux/features/auth/authSlice'
import { createAuction } from '../redux/features/auction/auctionSlice'
import { HeaderComponent } from '../components/HeaderComponent.jsx'

import type { GetProp, UploadFile, UploadProps } from 'antd';


import React, { useState } from 'react';
import { PlusOutlined } from '@ant-design/icons';
import {
  Button,
  DatePicker,
  Form,
  Input,
  InputNumber,
  Upload,
} from 'antd';

const { RangePicker } = DatePicker;
const { TextArea } = Input;


export const CreateAuctionPage: React.FC = () => {


  const dispatch = useDispatch()

const [form] = Form.useForm();
  const [fileList, setFileList] = useState([]);
    function callMe(values) {
        console.log("I was called")
        console.log(fileList);
        fileList.forEach((file) => {
            const reader = new FileReader();
            reader.onload = e => {
                console.log(e.target.result);
            }
            reader.readAsText(file);
        })
    }

const [uploading, setUploading] = useState(false);
    const props: UploadProps = {
        onRemove: (file) => {
            const index = fileList.indexOf(file);
            const newFileList = fileList.slice();
            newFileList.splice(index, 1);
            setFileList(newFileList);
        },
        beforeUpload: (file) => {
            setFileList([...fileList, file]);

            return false;
        },
        fileList,
        maxCount: 3,
        listType:"picture-card",
    };
    const onFinish = (values) => {
        console.log(fileList);
        try {
            dispatch(createAuction(values))
        } catch (error) {
            console.log(error)
        }
        console.log('Received values of form: ', values);
        form.resetFields()
    };
  return (
    <>
        <Form
            form={form}
            onFinish={onFinish}
            labelCol={{ span: 4 }}
            wrapperCol={{ span: 14 }}
            layout="horizontal"
        >
        <Form.Item
            name="title"
            label="Auction Title">
            <Input/>
        </Form.Item>
        <Form.Item 
            name="description"
            label="Description"
        >
            <TextArea rows={4}/>
        </Form.Item>
        <Form.Item 
            name="timeframe"
            label="Timeframe">
            <RangePicker />
        </Form.Item>
        <Form.Item 
            name="startPrice"
            label="Start Price">
          <InputNumber />
        </Form.Item>
        <Form.Item
            label="Lot Photo"
            name="files"
            initialValue={fileList}
        >
            <Upload
                {...props}
            >
                <button style={{ border: 0, background: 'none' }} type="button">
                <PlusOutlined />
                <div style={{ marginTop: 8 }}>Upload</div>
                </button>
            </Upload>
        </Form.Item>
        <Form.Item label="Button">
            <Button htmlType="submit">Submit</Button>
        </Form.Item>
      </Form>
    </>
  );
};

