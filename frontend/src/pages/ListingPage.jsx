import { React, useState, useEffect } from 'react'
import { Space, Table, Tag } from 'antd';
import type { TableProps } from 'antd';
import { useDispatch, useSelector } from 'react-redux'
import { fetchAllAuctions } from '../redux/features/auction/auctionSlice'

const columns: TableProps<DataType>['columns'] = [
  {
    title: 'Title',
    dataIndex: ['title', 'id' ],
    render: (title) => <a href={"/auction/"}  >{title}</a>,
  },
];

export const ListingPage: React.FC = () => {
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
        <Table
            loading={loading}
            columns={columns}
            dataSource={auctions}
            pagination={{
                pageSize: pageSize,
                total: totalPages,
            }}
        />
    )
}
