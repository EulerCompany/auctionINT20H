import { React, useEffect } from 'react'
import { Avatar, List } from 'antd';
import { useDispatch } from 'react-redux'
import { useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { fetchBetsForAuction } from '../redux/features/auction/auctionSlice';

export const BetHistoryComponent = () => {

    const bets = useSelector((state) => {
        console.log("bets:", state.auction.bets)
        return state.auction.bets
    })


  return (
    <div className='ml-10'>
    <h2>Bet history</h2>
    <List
    className='w-1/12'
    itemLayout="horizontal"
    dataSource={bets}
    renderItem={(item, index) => (
      <List.Item>
        <List.Item.Meta
          avatar={<Avatar src={`https://api.dicebear.com/7.x/miniavs/svg?seed=${index}`} />}
          title={<div>{item.user}</div>}
          description={<div>{item.bet}$</div>}
        />
      </List.Item>
    )}
  />
  </div>
  )
}
