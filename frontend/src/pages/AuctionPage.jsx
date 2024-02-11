import { React, useEffect } from 'react'
import { useParams } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { EditOutlined, EllipsisOutlined, SettingOutlined } from '@ant-design/icons';
import { Avatar, Button, Card, Image, Tag, InputNumber } from 'antd';
import { BetHistoryComponent } from '../components/BetHistoryComponent';
import { useDispatch } from 'react-redux'
import { fetchAllAuctions, fetchAuctionPhotos } from '../redux/features/auction/auctionSlice'
const { Meta } = Card;

export function AuctionPage () {
  const { id } = useParams();
  const dispatch = useDispatch()
    
  console.log("id is " + id);
  useEffect(() => {
    dispatch(fetchAllAuctions());
      if (id) {
          dispatch(fetchAuctionPhotos(id));
      }
  }, [dispatch, id])

  const isOwner = false

  const auction = useSelector(state =>
    state.auction.auctions.find(auction => auction.Id.toString() === id)
  );

  if (!auction) {
    console.log(id)
    return <div>Auction not found</div>;
  }

  return (
    <div className=' content-center flex justify-center'>
    <Card
    
    
    cover={
      <div className='flex'>
      <Image.PreviewGroup
      
    preview={{
      onChange: (current, prev) => console.log(`current index: ${current}, prev index: ${prev}`),
    }}
  >
    <Image width={200} src="https://gw.alipayobjects.com/zos/rmsportal/KDpgvguMpGfqaHPjicRK.svg" />
    <Image
      width={200}
      src="https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg"
    />
  </Image.PreviewGroup>
  </div>
    }
    actions={ isOwner && [
      <EditOutlined key="edit" />,
      <Button className='w-3/4' type="primary"  >Stop Auction</Button>,
    ] || !isOwner && [
      <InputNumber min={auction.CurrentPrice + 1} addonAfter="$" defaultValue={auction.CurrentPrice + 1} />,
      <Button className='w-3/4' type="primary"  >Make Bet</Button>,
    ]}
  >
    <Meta
      title={<div>
        <h2 className='text-3xl'>
        {auction.Title}
        <Tag className='ml-5' color={'green'}>
                {auction.Status}
            </Tag>

        </h2>
        <h3>DESCRIPTION:</h3>
        <p>{auction.Description}</p>
        <div className='flex justify-start'>
          <div className='m me-48'> 
        <h3>START PRICE:</h3>
        <p className='text-2xl text-red-600'>{auction.StartPrice} $</p>
        </div>
        <div>
        <h3>CURRENT PRICE:</h3>
        <p className='text-2xl text-red-600'>{auction.CurrentPrice} $</p>
        </div>
        </div>
        <div className='flex justify-start'>
        <div className='mr-20'>
        <h3>START DATE:</h3>
        <p>{new Date(auction.StartDate.Time).toUTCString()}</p>
        </div>
        <div>
        <h3>END DATE:</h3>
        <p>{new Date(auction.EndDate.Time).toUTCString()}</p>
        </div>
        </div>
        </div>}
      description={"ATTENTION: Author can stop auction in any time"}
    />
    
  </Card>
  <BetHistoryComponent/>
  </div>
  )
}
