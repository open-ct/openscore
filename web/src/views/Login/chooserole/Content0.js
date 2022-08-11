import React from "react";
import QueueAnim from "rc-queue-anim";
import {Col, Row} from "antd";
import {getChildrenToRender} from "../../../Util";

class Content extends React.PureComponent {
  render() {
    const {dataSource, ...props} = this.props;
    const {
      wrapper,
      titleWrapper,
      page,
      childWrapper,
    } = dataSource;
    return (
      <div {...props} {...wrapper}>
        <div {...page}>
          <div {...titleWrapper}>
            {titleWrapper.children.map(getChildrenToRender)}
          </div>
          <QueueAnim
            type="bottom"
            key="block"
            leaveReverse
            component={Row}
            componentProps={childWrapper}
          >
            {childWrapper.children.map((block, i) => {
              const {children: item, ...blockProps} = block;
              return (
                <Col key={i.toString()} {...blockProps}>
                  <div {...item}>
                    {item.children.map(getChildrenToRender)}
                  </div>
                </Col>
              );
            })}
          </QueueAnim>
        </div>
      </div>
    );
  }
}

export default Content;
