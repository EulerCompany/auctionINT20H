
import { LockOutlined, UserOutlined } from '@ant-design/icons'
import { NavLink, useNavigate } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { isAuth } from '../redux/features/auth/authSlice'
import { createAuction } from '../redux/features/auction/auctionSlice'
import { HeaderComponent } from '../components/HeaderComponent.jsx'
import { GetProp, UploadFile, UploadProps } from 'antd';
import React, { useState } from 'react';
import { PlusOutlined } from '@ant-design/icons';
import { Button, DatePicker, Form, Input, InputNumber, Upload } from 'antd';

const { RangePicker } = DatePicker;
const { TextArea } = Input;


export const CreateAuctionPage = () => {

    const formatDayjsObjectToISO = (dayjsObj) => dayjsObj.toISOString();

    const navigate = useNavigate();

    const dispatch = useDispatch()

    const [form] = Form.useForm();

    const [fileList, setFileList] = useState([]);

    const userId = useSelector((state) => {
        return state.auth.userId;
    });
        
    const [uploading, setUploading] = useState(false);

    const props = {
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

    const onFinish = async (values) => {
        if (isSubmitting) return;
        setIsSubmitting(true);
        const [startDayjs, endDayjs] = values.timeframe;
        const formattedStartDate = formatDayjsObjectToISO(startDayjs);
        const formattedEndDate = formatDayjsObjectToISO(endDayjs);
        delete values.timeframe;
        values.start_date = formattedStartDate;
        values.end_date = formattedEndDate;
        try {
        const base64Files = await Promise.all(fileList.map(file => {
            return new Promise((resolve, reject) => {
                const reader = new FileReader();
                reader.readAsDataURL(file);
                reader.onload = () => {
                const base64 = reader.result.split(',')[1];
                resolve({
                    base64,
                    name: file.name,
                    type: file.type,
                    size: file.size
                    // Add more metadata properties as needed
                });
            };
            reader.onerror = error => reject(error);            });
        }));

        values.files = base64Files;

        console.log(values);

        dispatch(createAuction(values))
        form.resetFields()
        } catch (error) {
            console.log(error)
        }
        console.log('Received values of form: ', values);
        form.resetFields()
        };

        const [isSubmitting, setIsSubmitting] = useState(false);

  return (
    <div>
        <Form
            form={form}
            onFinish={onFinish}
            labelCol={{ span: 4 }}
            wrapperCol={{ span: 14 }}
            layout="horizontal">
                
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

        <Form.Item 
            name="timeframe"
            label="Timeframe">
            <RangePicker />
        </Form.Item>

        <Form.Item 
            name="start_price"
            label="Start Price">
          <InputNumber />
        </Form.Item>

        <Form.Item
            label="Lot Photo"
            name="files"
            initialValue={fileList}>

            <Upload {...props} >
                <button style={{ border: 0, background: 'none' }} type="button">
                <PlusOutlined />
                <div style={{ marginTop: 8 }}>Upload</div>
                </button>
            </Upload>
        </Form.Item>

        <Form.Item label="Button">
            <Button type='primary' htmlType="submit" disabled={isSubmitting}>Submit</Button>
        </Form.Item>

      </Form>
    </div>
  );
};

