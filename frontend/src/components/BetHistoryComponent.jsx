import React from 'react'
import { Avatar, List } from 'antd';
const data = [
  {
    title: 'Design1',
    bet: "100",
  },
  {
    title: 'Design2',
    bet: "100",
  },
  {
    title: 'Design3',
    bet: "100",
  },
  {
    title: 'Design4',
    bet: "100",
  },
];

export const BetHistoryComponent = () => {
  return (
    <div className='ml-10'>
    <h2>Bet history</h2>
    <List
    className='w-1/12'
    itemLayout="horizontal"
    dataSource={data}
    renderItem={(item, index) => (
      <List.Item>
        <List.Item.Meta
          avatar={<Avatar src={`https://api.dicebear.com/7.x/miniavs/svg?seed=${index}`} />}
          title={<div>{item.title}</div>}
          description={<div>{item.bet}$</div>}
        />
      </List.Item>
    )}
  />
  </div>
  )
}
