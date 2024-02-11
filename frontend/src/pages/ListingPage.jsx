import { React, useState, useEffect } from 'react'
import { Space, Table, Tag } from 'antd';
import { TableProps } from 'antd';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux'
import { fetchAllAuctions } from '../redux/features/auction/auctionSlice'



const columns = [
  {
    title: 'Title',
    dataIndex: ['Title'],
    key: 'Title',
    render: (text, record) => (
        <NavLink to={`/auction/${record.Id}`}> {text}</NavLink>        
      ),
  },
  {
    title: 'Current Price',
    dataIndex: 'CurrentPrice',
    key: 'CurrentPrice',
    render: ( CurrentPrice ) => (
        <>
            <Tag className='s text-lg' color={'volcano'} key={CurrentPrice}>
                {CurrentPrice} $
            </Tag>
        </>
      ),
  },
  {
    title: 'End date',
    dataIndex: 'EndDate',
    key: 'EndDate',
    render: ( EndDate ) => (
        <>
            {EndDate ? new Date(EndDate.Time).toUTCString() : 'N/A'}
        </>
      ),
    },
  {
    title: 'Status',
    dataIndex: 'Status',
    key: 'Status',
    render: ( Status ) => (
        <>
            <Tag color={'green'} key={Status}>
                {Status.toUpperCase()}
            </Tag>
        </>
      ),
    },
  
];

export const ListingPage = () => {
    

    const loading = useSelector((state) => {
        console.log(state);
        return state.auction.loading
    })

    const auctions = useSelector((state) => {
        return state.auction.auctions
    })

    const pageSize = useSelector((state) => {
        return state.auction.pageSize
    })
    
    const totalPages = useSelector((state) => {
        return state.auction.totalPages
    })

    const dispatch = useDispatch()

    useEffect(() => {
        dispatch(fetchAllAuctions());
    }, [dispatch])


    return (
        <div>
        <h2>Available auctions</h2>
        <Table
            loading={loading}
            columns={columns}
            dataSource={auctions}
            pagination={{
                pageSize: pageSize,
                total: totalPages,
            }}
        />
        </div>
    )
}
