import { React, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import { NavLink, useLocation, useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { EditOutlined, EllipsisOutlined, SettingOutlined } from '@ant-design/icons';
import { Avatar, Button, Card, Image, Tag, InputNumber, List } from 'antd';
import { BetHistoryComponent } from '../components/BetHistoryComponent';
import { useDispatch } from 'react-redux'
import { fetchAllAuctions, fetchAuctionPhotos, fetchBetsForAuction, makeBet  } from '../redux/features/auction/auctionSlice'
const { Meta } = Card;

export function AuctionPage() {
  const { id } = useParams();
  const navigate = useNavigate();

  const dispatch = useDispatch()
  const [bet_value, setBet_value] = useState(0)

  console.log("id is " + id);
  useEffect(() => {
    dispatch(fetchAllAuctions());
    if (id) {
      dispatch(fetchBetsForAuction(id))
      dispatch(fetchAuctionPhotos(id));
      
    }
  }, [dispatch, id])


  const auction = useSelector(state =>
    state.auction.auctions.find(auction => auction.Id.toString() === id)
  );

   



  const isOwner = useSelector((state) => {
    console.log(state);
    if(auction) {
    return state.auth.userId === auction.AuthorId
    }
    
  })

  const photos = useSelector(state => {
    return state.auction.photos;
  });

  const bets = useSelector((state) => {
    console.log("bets:", state.auction.bets)
    return state.auction.bets
})




  if (!auction) {
    console.log(id)
    return <div>Auction not found</div>;
  }

  

  const onChange = (value) => {
    setBet_value(value)
    console.log('changed', value);
  };

  const make_bet = (bet_v) => {
    console.log("iddddd", {'bet': bet_v})
    try {
      dispatch(makeBet({"id": id, "bet": bet_v}))
    } catch (error) {
      console.log(error)
    }
    console.log('Received values of form: ', bet_v);
    window.location.reload(false);
  };
  console.log(photos);

  const editHandler = () => {
    navigate(`/auction/${id}/edit`)
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
              {photos.map((photo, index) => (
                <Image
                  key={index}
                  width={200}
                  src={photo}
                />
              ))}
            </Image.PreviewGroup>
          </div>
        }
        actions={isOwner && [
          <EditOutlined onClick={editHandler} key="edit" />,
          <Button className='w-3/4' type="primary"  >Stop Auction</Button>,
        ] || ((!isOwner && auction) && [
          <InputNumber min={auction.CurrentPrice + 1} onChange={onChange} addonAfter="$"  />,
          <Button className='w-3/4' type="primary" onClick={() => {make_bet(bet_value)}} >Make Bet</Button>,
        ])}
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
      {/* <BetHistoryComponent  /> */}
      <div className='ml-10'>
    <h2>Bet history</h2>
    { bets && <List
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
  />}
  </div>
    </div>
  )
}
