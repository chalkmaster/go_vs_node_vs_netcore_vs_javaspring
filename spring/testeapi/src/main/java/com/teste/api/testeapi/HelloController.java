package com.teste.api.testeapi;

import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;


@RestController
public class HelloController {

    private static HashMap<Integer, Integer> database = new HashMap<Integer, Integer>();
    @RequestMapping(value="/player/progress/{id}", method=RequestMethod.GET)
    public Resp get(@PathVariable int id) {
        if (!database.containsKey(id)) {
            database.put(id, 0);
        }
        return new Resp(database.get(id));
    }
    
    @RequestMapping(value="/player/progress/{id}", method=RequestMethod.PATCH)
    public void patch(@PathVariable int id) {
        if (!database.containsKey(id)) {
            database.put(id, 0);
        }
        database.replace(id,database.get(id) + 1);
    }
}