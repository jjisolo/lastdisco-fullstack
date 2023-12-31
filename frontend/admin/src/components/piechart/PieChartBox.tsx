import './PieChartBox.scss'
import {Cell, Pie, PieChart, ResponsiveContainer, Tooltip} from "recharts";

const data = [
    { name: "Животные", value: 400, color: "#0088FE" },
    { name: "Растения", value: 300, color: "#00C49F" },
    { name: "Фэндом", value: 300, color: "#FFBB28" },
    { name: "Разное", value: 200, color: "#FF8042" },
];

const PieChartBox = () => {
    return (
        <div className="pieChart">
           <h1>
                Популярные типы изделий
           </h1>
           <div className="chart">
               <ResponsiveContainer width="99%" height={300}>
                   <PieChart>
                       <Tooltip
                           contentStyle={{ background:"white", borderRadius:"5px", }}
                       />
                       <Pie
                           data={data}
                           innerRadius="70%"
                           outerRadius="90%"
                           paddingAngle={5}
                           dataKey="value"
                       >
                           {data.map((item) => (
                               <Cell
                                   key={item.name}
                                   fill={item.color}
                               />
                           ))}
                       </Pie>
                   </PieChart>
               </ResponsiveContainer>
           </div>
           <div className="options">
               {data.map((item) => (
                 <div className="option" key={item.name}>
                     <div className="title">
                         <div className="dot" style={{ backgroundColor: item.color }} />
                         <div className="dotTitle">
                             {item.name}
                         </div>
                     </div>
                     <span className="dotDescription">
                        {item.value}
                     </span>
                 </div>
               ))}
           </div>
        </div>
    )
}

export default PieChartBox