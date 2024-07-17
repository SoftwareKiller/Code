package org.jayden.sprintbootmybits.domain;

public class App {
    private int id;
    private String username;
    private String app_id;

    public String getApp_id() {
        return app_id;
    }

    public void setApp_id(String app_id) {
        this.app_id = app_id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    @Override
    public String toString() {
        return "App{" +
                "id=" + id +
                ", username='" + username + '\'' +
                ", app_id='" + app_id + '\'' +
                '}';
    }
}
