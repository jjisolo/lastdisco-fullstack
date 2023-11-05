import './ChartBox.scss'
import {Link} from "react-router-dom";
import {Line, LineChart, ResponsiveContainer, Tooltip} from "recharts";

type Props = {
    color: string;
    icon: string;
    title: string;
    dataKey: string;
    number: number | string;
    percentage: number;
    chartData: object[];
}

const ChartBox = (props: Props) => {
    return (
       <div className="chartBox">
           <div className="boxInfo">
               <div className="title">
                   <img src={props.icon} alt="" />
                   <span>
                       { props.title }
                   </span>
               </div>
                   <h1>
                       { props.number }
                   </h1>
                   <Link to="/" style={{ color: "#8884d8", fontSize: "18px" }}>
                       Больше
                   </Link>
           </div>
           <div className="chartInfo">
               <div className="chart">
                   <ResponsiveContainer width="100%" height="100%">
                       <LineChart data={props.chartData}>
                           <Tooltip
                            contentStyle={{background: "transparent", border:"none"}}
                            labelStyle={{display:"none"}}
                            position={{x:27,y:65}}
                           />
                           <Line
                               type="monotone"
                               dataKey={props.dataKey}
                               stroke={props.color}
                               strokeWidth={2}
                               dot={false}
                           />
                       </LineChart>
                   </ResponsiveContainer>
               </div>
               <div className="texts">
                   <span className="percentage" style={{ color: props.percentage < 0 ? "tomato": "limegreen"}}>
                       { props.percentage }%
                   </span>
                   <span className="duration">
                        За этот месяц
                   </span>
               </div>
           </div>
       </div>
    )
}

export default ChartBox