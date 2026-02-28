package com.example.kafka.producer_app.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Service;

import com.example.kafka.producer_app.producer.Course;

@Service
public class KafkaService {
    @Autowired
    private KafkaTemplate<String, Course> kafkaTemplate;
    
    public String sendMessage(Course course) {
        kafkaTemplate.send("test-topic-4", "course", course);
        return "Course message sent to kakfa server";
    }
}
