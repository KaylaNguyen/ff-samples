package com.example;

import com.google.cloud.functions.CloudEventsFunction;
import io.cloudevents.CloudEvent;

public class CEFunction implements CloudEventsFunction {
  @Override
  public void accept(CloudEvent event) throws Exception {
    // throw new Exception("Intended Exception");
  }
}